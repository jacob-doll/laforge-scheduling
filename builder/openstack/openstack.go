package openstack

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	"golang.org/x/sync/semaphore"

	"github.com/gen0cide/laforge/ent"
	"github.com/gen0cide/laforge/logging"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/secgroups"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
)

const (
	ID          = "openstack"
	Name        = "Openstack"
	Description = "Builder that interfaces with Openstack"
	Author      = "Tenchi Mata <github.com/0xk7>"
	Version     = "0.1"
)

type OpenstackBuilder struct {
	Config             OpenstackBuilderConfig
	HttpClient         http.Client
	Logger             *logging.Logger
	DeployWorkerPool   *semaphore.Weighted
	TeardownWorkerPool *semaphore.Weighted
}

type OpenstackBuilderConfig struct {
	AuthUrl            string            `json:"auth_url"`
	ServerUrl          string            `json:"laforge_server_url"`
	IdentityVersion    string            `json:"identify_version"`
	Username           string            `json:"username"`
	Password           string            `json:"password"`
	ProjectID          string            `json:"project_id"`
	ProjectName        string            `json:"project_name"`
	RegionName         string            `json:"region_name"`
	DomainName         string            `json:"domain_name"`
	DomainId           string            `json:"domain_id"`
	MaxBuildWorkers    int               `json:"max_build_workers"`
	MaxTeardownWorkers int               `json:"max_teardown_workers"`
	Flavors            map[string]string `json:"flavors"`
	Images             map[string]string `json:"images"`
	ExternalNetworkId  string            `json:"external_network_id"`
}

func (builder OpenstackBuilder) ID() string {
	return ID
}

func (builder OpenstackBuilder) Name() string {
	return Name
}

func (builder OpenstackBuilder) Description() string {
	return Description
}

func (builder OpenstackBuilder) Author() string {
	return Author
}

func (builder OpenstackBuilder) Version() string {
	return Version
}

func (builder OpenstackBuilder) generateBuildID(build *ent.Build) string {
	buildId, err := build.ID.MarshalText()
	if err != nil {
		buildId = []byte(fmt.Sprint(build.Revision))
	}
	return fmt.Sprintf("%s", buildId)
}

func (builder OpenstackBuilder) generateVmName(competition *ent.Competition, team *ent.Team, host *ent.Host, build *ent.Build) string {
	return (competition.HclID + "-Team-" + fmt.Sprintf("%02d", team.TeamNumber) + "-" + host.Hostname + "-" + builder.generateBuildID(build))
}

func (builder OpenstackBuilder) generateRouterName(competition *ent.Competition, team *ent.Team, build *ent.Build) string {
	return (competition.HclID + "-Team-" + fmt.Sprintf("%02d", team.TeamNumber) + "-" + builder.generateBuildID(build))
}

func (builder OpenstackBuilder) generateNetworkName(competition *ent.Competition, team *ent.Team, network *ent.Network, build *ent.Build) string {
	return (competition.HclID + "-Team-" + fmt.Sprintf("%02d", team.TeamNumber) + "-" + network.Name + "-" + builder.generateBuildID(build))
}

func (builder OpenstackBuilder) newAuthProvider() (*gophercloud.ProviderClient, error) {
	u, err := url.Parse(builder.Config.AuthUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to parse auth_url \"%s\" from builder config", builder.Config.AuthUrl)
	}
	u.Path = path.Join(u.Path, builder.Config.IdentityVersion)

	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: u.String(),
		Username:         builder.Config.Username,
		Password:         builder.Config.Password,
		TenantID:         builder.Config.ProjectID,
		TenantName:       builder.Config.ProjectName,
	}
	if builder.Config.DomainName != "" {
		authOpts.DomainName = builder.Config.DomainName
	} else {
		authOpts.DomainID = builder.Config.DomainId
	}
	return openstack.AuthenticatedClient(authOpts)
}

func (builder OpenstackBuilder) DeployHost(ctx context.Context, entProvisionedHost *ent.ProvisionedHost) (err error) {
	entHost, err := entProvisionedHost.QueryProvisionedHostToHost().Only(ctx)
	if err != nil {
		return err
	}

	provider, err := builder.newAuthProvider()
	if err != nil {
		return fmt.Errorf("failed to authenticate: %v", err)
	}
	endpointOpts := gophercloud.EndpointOpts{
		Region: builder.Config.RegionName,
	}

	computeClient, err := openstack.NewComputeV2(provider, endpointOpts)
	if err != nil {
		return fmt.Errorf("failed to create compute v2 client: %v", err)
	}

	entBuild, err := entProvisionedHost.QueryProvisionedHostToPlan().QueryPlanToBuild().Only(ctx)
	if err != nil {
		return err
	}

	entCompetition, err := entBuild.QueryBuildToCompetition().Only(ctx)
	if err != nil {
		return err
	}

	// network, err := provisionedHost.QueryProvisionedHostToProvisionedNetwork().QueryProvisionedNetworkToNetwork().Only(ctx)
	// if err != nil {
	// 	return err
	// }

	entTeam, err := entProvisionedHost.QueryProvisionedHostToProvisionedNetwork().QueryProvisionedNetworkToTeam().Only(ctx)
	if err != nil {
		return err
	}

	// generate vm name from ent
	vmName := builder.generateVmName(entCompetition, entTeam, entHost, entBuild)
	// networkName := builder.generateNetworkName(competition, team, network, build)

	entProvisionedNetwork, err := entProvisionedHost.QueryProvisionedHostToProvisionedNetwork().Only(ctx)
	if err != nil {
		return fmt.Errorf("failed to query provisioned network from provisioned host: %v", err)
	}

	networkAddrParts := strings.Split(entProvisionedNetwork.Cidr, "/")
	networkAddr := networkAddrParts[0]
	networkOctetStrings := strings.Split(networkAddr, ".")
	networkOctets := []byte{0, 0, 0, 0}
	for i, octetString := range networkOctetStrings {
		octet, err := strconv.Atoi(octetString)
		if err != nil {
			return fmt.Errorf("error while parsing IPv4 Address %s: %v", entProvisionedNetwork.Cidr, err)
		}
		networkOctets[i] = byte(octet)
	}

	_, ipv4Net, err := net.ParseCIDR(entProvisionedNetwork.Cidr)
	if err != nil {
		return fmt.Errorf("error while parsing cidr: %v", err)
	}
	if len(ipv4Net.Mask) != 4 {
		return fmt.Errorf("mask is not correct length")
	}
	hostAddress := strings.Join(append(networkOctetStrings[:3], fmt.Sprint(entHost.LastOctet)), ".")

	err = builder.DeployWorkerPool.Acquire(ctx, int64(1))
	if err != nil {
		return
	}
	defer builder.DeployWorkerPool.Release(int64(1))

	builder.Logger.Log.Debugf("Creating security group %s", vmName)

	opts := secgroups.CreateOpts{
		Name:        vmName,
		Description: fmt.Sprintf("Sec group for VM %s", vmName),
	}
	secGroup, err := secgroups.Create(computeClient, opts).Extract()
	if err != nil {
		return fmt.Errorf("failed to create security group: %v", err)
	}
	entProvisionedHost.Vars["openstack_secgroup_id"] = secGroup.ID
	err = entProvisionedHost.Update().SetVars(entProvisionedHost.Vars).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update provisioned host vars: %v", err)
	}

	fmt.Printf("%v\n%v", entHost.ExposedTCPPorts, entHost.ExposedUDPPorts)

	ruleOpts := make([]secgroups.CreateRuleOpts, 0)

	for _, port := range entHost.ExposedTCPPorts {
		destPort, err := strconv.Atoi(port)
		if err != nil {
			return fmt.Errorf("could not convert TCP port \"%s\" to integer: %v", port, err)
		}
		opts := secgroups.CreateRuleOpts{
			ParentGroupID: secGroup.ID,
			FromPort:      destPort,
			ToPort:        destPort,
			IPProtocol:    "TCP",
			CIDR:          "0.0.0.0/0",
		}
		ruleOpts = append(ruleOpts, opts)
	}

	for _, port := range entHost.ExposedUDPPorts {
		destPort, err := strconv.Atoi(port)
		if err != nil {
			return fmt.Errorf("could not convert UDP port \"%s\" to integer: %v", port, err)
		}
		opts := secgroups.CreateRuleOpts{
			ParentGroupID: secGroup.ID,
			FromPort:      destPort,
			ToPort:        destPort,
			IPProtocol:    "UDP",
			CIDR:          "0.0.0.0/0",
		}
		ruleOpts = append(ruleOpts, opts)
	}

	for _, opts := range ruleOpts {
		fmt.Printf("%v", ruleOpts)
		_, err = secgroups.CreateRule(computeClient, opts).Extract()
		if err != nil {
			return fmt.Errorf("failed to create security group %s rule to port %d", opts.IPProtocol, opts.FromPort)
		}
	}

	builder.Logger.Log.Debugf("Deploying host with image \"%s\" and flavor \"%s\"", builder.Config.Images[entHost.OS], builder.Config.Flavors[entHost.InstanceSize])

	var adminPassword string
	if len(entHost.OverridePassword) > 0 {
		adminPassword = entHost.OverridePassword
	} else {
		adminPassword = entCompetition.RootPassword
	}
	agentFile, err := entProvisionedHost.QueryProvisionedHostToGinFileMiddleware().First(ctx)
	if err != nil {
		return fmt.Errorf("failed to query gin file middleware from provisioned host: %v", err)
	}
	agentUrl := fmt.Sprintf("%s/api/download/%s", builder.Config.ServerUrl, agentFile.URLID)

	var userData string
	if strings.HasPrefix(entHost.OS, "w2k") {
		userData = fmt.Sprintf(`<script>
powershell -Command mkdir $env:PROGRAMDATA\Laforge -Force
powershell -Command do{	$test = Test-Connection 1.1.1.1 -Quiet; Start-Sleep -s 5}until($test)
powershell -Command Invoke-WebRequest %s -OutFile $env:PROGRAMDATA\Laforge\laforge.exe
powershell -Command %%PROGRAMDATA%%\Laforge\laforge.exe -service install
powershell -Command %%PROGRAMDATA%%\Laforge\laforge.exe -service start
powershell -Command logoff
</script>`, agentUrl)
	} else {
		userData = fmt.Sprintf(`#!/bin/bash
while [ ! -f "/laforge.bin" ]
do
curl -sL -o /laforge.bin %s
sleep 10
done
chmod +x /laforge.bin
cd /
./laforge.bin -service install
./laforge.bin -service start
`, agentUrl)
	}
	//build server
	server, err := servers.Create(computeClient, servers.CreateOpts{
		Name:           vmName,
		ImageRef:       builder.Config.Images[entHost.OS],
		FlavorRef:      builder.Config.Flavors[entHost.InstanceSize],
		SecurityGroups: []string{entProvisionedHost.Vars["openstack_secgroup_id"]}, // TODO: Remove this and make it dynamic
		UserData:       []byte(userData),
		Networks: []servers.Network{{
			UUID:    entProvisionedNetwork.Vars["openstack_network_id"],
			FixedIP: hostAddress,
		}},
		AdminPass:  adminPassword,
		AccessIPv4: hostAddress,
	}).Extract()
	if err != nil {
		builder.Logger.Log.Errorf("Unable to create server: %v", err)
		return
	}
	builder.Logger.Log.WithField("root_password", server.AdminPass).Debugf("Server ID: %s", server.ID)

	id := server.ID
	newVars := entHost.Vars
	newVars["openstack_instance_id"] = id
	err = entProvisionedHost.Update().SetVars(newVars).Exec(ctx)
	if err != nil {
		return err
	}
	builder.Logger.Log.Infof("Created tagged instance with ID " + id)

	return
}

func (builder OpenstackBuilder) DeployNetwork(ctx context.Context, entProvisionedNetwork *ent.ProvisionedNetwork) (err error) {
	entBuild, err := entProvisionedNetwork.QueryProvisionedNetworkToBuild().Only(ctx)
	if err != nil {
		return fmt.Errorf("couldn't query build from network \"%s\": %v", entProvisionedNetwork.Name, err)
	}
	entEnvironment, err := entProvisionedNetwork.QueryProvisionedNetworkToBuild().QueryBuildToEnvironment().Only(ctx)
	if err != nil {
		return fmt.Errorf("couldn't query build from network \"%s\": %v", entProvisionedNetwork.Name, err)
	}
	entCompetition, err := entProvisionedNetwork.QueryProvisionedNetworkToBuild().QueryBuildToEnvironment().QueryEnvironmentToCompetition().Only(ctx)
	if err != nil {
		return fmt.Errorf("couldn't query build from environment \"%s\": %v", entEnvironment.Name, err)
	}
	entNetwork, err := entProvisionedNetwork.QueryProvisionedNetworkToNetwork().Only(ctx)
	if err != nil {
		return fmt.Errorf("couldn't query build from network \"%s\": %v", entProvisionedNetwork.Name, err)
	}
	entTeam, err := entProvisionedNetwork.QueryProvisionedNetworkToTeam().Only(ctx)
	if err != nil {
		return fmt.Errorf("couldn't query build from network \"%s\": %v", entProvisionedNetwork.Name, err)
	}

	//add configuration here
	provider, err := builder.newAuthProvider()
	if err != nil {
		return fmt.Errorf("failed to authenticate: %v", err)
	}
	endpointOpts := gophercloud.EndpointOpts{
		Name:   "neutron",
		Region: builder.Config.RegionName,
	}

	client, err := openstack.NewNetworkV2(provider, endpointOpts)
	if err != nil {
		return err
	}

	osNetwork, err := networks.Create(client, networks.CreateOpts{
		Name:         builder.generateNetworkName(entCompetition, entTeam, entNetwork, entBuild),
		AdminStateUp: gophercloud.Enabled,
	}).Extract()
	if err != nil {
		return fmt.Errorf("failed to create network: %v", err)
	}
	newVars := entNetwork.Vars
	newVars["openstack_network_id"] = osNetwork.ID
	err = entProvisionedNetwork.Update().SetVars(newVars).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update provisioned network vars: %v", err)
	}

	networkAddrParts := strings.Split(entProvisionedNetwork.Cidr, "/")
	networkAddr := networkAddrParts[0]
	networkOctetStrings := strings.Split(networkAddr, ".")
	networkOctets := []byte{0, 0, 0, 0}
	for i, octetString := range networkOctetStrings {
		octet, err := strconv.Atoi(octetString)
		if err != nil {
			return fmt.Errorf("error while parsing IPv4 Address %s: %v", entProvisionedNetwork.Cidr, err)
		}
		networkOctets[i] = byte(octet)
	}

	_, ipv4Net, err := net.ParseCIDR(entProvisionedNetwork.Cidr)
	if err != nil {
		return fmt.Errorf("error while parsing cidr: %v", err)
	}
	if len(ipv4Net.Mask) != 4 {
		return fmt.Errorf("mask is not correct length")
	}
	routerAddress := strings.Join(append(networkOctetStrings[:3], "254"), ".")

	osSubnet, err := subnets.Create(client, subnets.CreateOpts{
		NetworkID:       osNetwork.ID,
		CIDR:            entNetwork.Cidr,
		Name:            builder.generateNetworkName(entCompetition, entTeam, entNetwork, entBuild),
		Description:     fmt.Sprintf("%s@%d Subnet for  Network \"%s\" for Team %d", entEnvironment.Name, entBuild.EnvironmentRevision, entNetwork.Name, entTeam.TeamNumber),
		AllocationPools: []subnets.AllocationPool{},
		GatewayIP:       &routerAddress,
		IPVersion:       gophercloud.IPv4,
	}).Extract()
	if err != nil {
		return fmt.Errorf("failed to create subnet: %v", err)
	}
	newVars["openstack_subnet_id"] = osSubnet.ID
	err = entProvisionedNetwork.Update().SetVars(newVars).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update provisioned network vars: %v", err)
	}

	osPort, err := ports.Create(client, ports.CreateOpts{
		NetworkID:    osNetwork.ID,
		Description:  fmt.Sprintf("%s@%d Router Interface on Network \"%s\" for Team %d", entEnvironment.Name, entBuild.EnvironmentRevision, entNetwork.Name, entTeam.TeamNumber),
		AdminStateUp: gophercloud.Enabled,
		FixedIPs: []ports.IP{{
			SubnetID:  osSubnet.ID,
			IPAddress: routerAddress,
		}},
	}).Extract()
	if err != nil {
		return fmt.Errorf("failed to create port for router: %v", err)
	}
	newVars["openstack_subnet_port_id"] = osPort.ID
	err = entProvisionedNetwork.Update().SetVars(newVars).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update provisioned network vars: %v", err)
	}

	osRouterId, exists := entTeam.Vars["openstack_router_id"]
	if !exists {
		return fmt.Errorf("failed to get openstack_router_id from team vars")
	}
	osInterface, err := routers.AddInterface(client, osRouterId, routers.AddInterfaceOpts{
		PortID: osPort.ID,
	}).Extract()
	if err != nil {
		return fmt.Errorf("failed to create router interface: %v", err)
	}
	newVars["openstack_router_interface_id"] = osInterface.ID
	err = entProvisionedNetwork.Update().SetVars(newVars).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update provisioned network vars: %v", err)
	}

	return
}

func (builder OpenstackBuilder) DeployTeam(ctx context.Context, entTeam *ent.Team) (err error) {
	entBuild, err := entTeam.QueryTeamToBuild().Only(ctx)
	if err != nil {
		return fmt.Errorf("couldn't query build from team: %v", err)
	}
	entEnvironment, err := entTeam.QueryTeamToBuild().QueryBuildToEnvironment().Only(ctx)
	if err != nil {
		return fmt.Errorf("couldn't query environment from team: %v", err)
	}
	entCompetition, err := entTeam.QueryTeamToBuild().QueryBuildToEnvironment().QueryEnvironmentToCompetition().Only(ctx)
	if err != nil {
		return fmt.Errorf("couldn't query competition from team: %v", err)
	}

	provider, err := builder.newAuthProvider()
	if err != nil {
		return fmt.Errorf("failed to authenticate: %v", err)
	}
	endpointOpts := gophercloud.EndpointOpts{
		Name:   "neutron",
		Region: builder.Config.RegionName,
	}

	client, err := openstack.NewNetworkV2(provider, endpointOpts)
	if err != nil {
		return err
	}

	osRouter, err := routers.Create(client, routers.CreateOpts{
		Name:         builder.generateRouterName(entCompetition, entTeam, entBuild),
		Description:  fmt.Sprintf("%s@%d Router for Team %d", entEnvironment.Name, entBuild.EnvironmentRevision, entTeam.TeamNumber),
		AdminStateUp: gophercloud.Enabled,
		GatewayInfo: &routers.GatewayInfo{
			NetworkID: builder.Config.ExternalNetworkId,
		},
	}).Extract()
	if err != nil {
		return fmt.Errorf("failed to create router: %v", err)
	}
	newVars := entTeam.Vars
	newVars["openstack_router_id"] = osRouter.ID
	err = entTeam.Update().SetVars(newVars).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update team vars: %v", err)
	}
	return
}

func (builder OpenstackBuilder) TeardownHost(ctx context.Context, provisionedHost *ent.ProvisionedHost) (err error) {
	// func (builder OpenstackBuilder) TeardownHost(ctx contecontext)
	// host, err := provisionedHost.QueryProvisionedHostToHost().Only(ctx)
	// if err != nil {
	// 	return err
	// }
	// entBuild, err := provisionedHost.QueryProvisionedHostToProvisionedNetwork().QueryProvisionedNetworkToBuild().Only(ctx)
	// if err != nil {
	// 	return err
	// }
	// entCompetition, err := entBuild.QueryBuildToCompetition().Only(ctx)
	// if err != nil {
	// 	return fmt.Errorf("couldn't query competition from build \"%s\": %v", entBuild.ID, err)
	// }
	// entTeam, err := provisionedHost.QueryProvisionedHostToProvisionedNetwork().QueryProvisionedNetworkToTeam().Only(ctx)
	// if err != nil {
	// 	return fmt.Errorf("couldn't query build from network \"%s\": %v", provisionedHost.ID, err)
	// }
	authOpts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return err
	}
	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		return err
	}
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	if err != nil {
		return err
	}

	// vmName := builder.generateVmName(entCompetition, entTeam, host, entBuild)

	err = builder.TeardownWorkerPool.Acquire(ctx, int64(1))
	if err != nil {
		return
	}
	defer builder.TeardownWorkerPool.Release(int64(1))

	// var instances []string
	instanceId := provisionedHost.Vars["openstack_instance_id"]
	result := servers.Delete(client, instanceId)
	fmt.Println(result)

	return
}

func (builder OpenstackBuilder) TeardownNetwork(ctx context.Context, provisionedNetwork *ent.ProvisionedNetwork) (err error) {
	// entBuild, err := provisionedNetwork.QueryProvisionedNetworkToBuild().Only(ctx)
	// if err != nil {
	// 	return fmt.Errorf("couldn't query build from network \"%s\": %v", provisionedNetwork.Name, err)
	// }
	// entEnvironment, err := provisionedNetwork.QueryProvisionedNetworkToBuild().QueryBuildToEnvironment().Only(ctx)
	// if err != nil {
	// 	return fmt.Errorf("couldn't query build from network \"%s\": %v", provisionedNetwork.Name, err)
	// }
	// entCompetition, err := entEnvironment.EnvironmentToCompetition(ctx)
	// if err != nil {
	// 	return fmt.Errorf("couldn't query build from environment \"%s\": %v", entEnvironment.Name, err)
	// }
	// entNetwork, err := provisionedNetwork.QueryProvisionedNetworkToNetwork().Only(ctx)
	// if err != nil {
	// 	return fmt.Errorf("couldn't query build from network \"%s\": %v", provisionedNetwork.Name, err)
	// }
	// entTeam, err := provisionedNetwork.QueryProvisionedNetworkToTeam().Only(ctx)
	// if err != nil {
	// 	return fmt.Errorf("couldn't query build from network \"%s\": %v", provisionedNetwork.Name, err)
	// }

	authOpts, err := openstack.AuthOptionsFromEnv()
	if err != nil {
		return err
	}
	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		return err
	}
	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	if err != nil {
		return err
	}

	// networkName := builder.generateNetworkName(entCompetition[0], entTeam, entNetwork, entBuild)

	err = builder.TeardownWorkerPool.Acquire(ctx, int64(1))
	if err != nil {
		return
	}
	defer builder.TeardownWorkerPool.Release(int64(1))

	networkId := provisionedNetwork.Vars["openstack_instance_id"]
	result := networks.Delete(client, networkId)
	fmt.Println(result)

	return
}

func (builder OpenstackBuilder) TeardownTeam(ctx context.Context, entTeam *ent.Team) (err error) {
	return
}
