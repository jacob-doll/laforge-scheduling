import { ChangeDetectorRef, Component, OnInit, OnDestroy } from '@angular/core';
import { MatSelectChange } from '@angular/material/select';
import { Subscription } from 'rxjs';
import { updateAgentStatuses } from 'src/app/models/agent.model';
import { AgentStatusQueryResult } from 'src/app/models/api.model';
import { Environment, resolveStatuses } from 'src/app/models/environment.model';
import { ApiService } from 'src/app/services/api/api.service';
import { SubheaderService } from 'src/app/_metronic/partials/layout/subheader/_services/subheader.service';

@Component({
  selector: 'app-manage',
  templateUrl: './monitor.component.html',
  styleUrls: ['./monitor.component.scss']
})
export class MonitorComponent implements OnInit, OnDestroy {
  // corpNetwork: ProvisionedNetwork = corp_network_provisioned;
  environment: Environment = null;
  envLoaded = false;
  environmentDetailsCols: string[] = ['TeamCount', 'AdminCIDRs', 'ExposedVDIPorts', 'maintainer'];
  agentPollingInterval: NodeJS.Timeout;
  pollingInterval = 60;
  loading = false;
  intervalOptions = [10, 30, 60, 120];

  constructor(private api: ApiService, private cdRef: ChangeDetectorRef, private subheader: SubheaderService) {
    this.subheader.setTitle('Monitor Agents');
    this.subheader.setDescription('View live data being sent from the host agents');
  }

  ngOnInit(): void {
    this.api.pullEnvironment('a3f73ee0-da71-4aa6-9280-18ad1a1a8d16').then(
      (env: Environment) => {
        this.environment = resolveStatuses(env);
        this.envLoaded = true;
        this.cdRef.detectChanges();
        this.initAgentStatusPolling();
      },
      (err) => {
        console.error('yep, cant connect');
        // console.error(err);
      }
    );
  }

  ngOnDestroy(): void {
    clearInterval(this.agentPollingInterval);
  }

  initAgentStatusPolling(): void {
    // Go ahead and query the statuses for the first time
    this.fetchAgentStatuses();
    // Set up the query to be polled every interval
    this.agentPollingInterval = setInterval(() => this.fetchAgentStatuses(), this.pollingInterval * 1000);
  }

  fetchAgentStatuses(): void {
    this.loading = true;
    this.cdRef.detectChanges();
    this.api.getAgentStatuses(this.environment.id).then((result: AgentStatusQueryResult) => {
      this.loading = false;
      this.environment = updateAgentStatuses(this.environment, result);
      this.cdRef.detectChanges();
    }, console.error);
  }
  // onBranchSelect(changeEvent: MatSelectChange) {
  onIntervalChange(changeEvent: MatSelectChange): void {
    // Update the interval based on select's value
    this.pollingInterval = changeEvent.value;
    // Stop the old polling
    clearInterval(this.agentPollingInterval);
    // Set up polling again with new interval
    this.agentPollingInterval = setInterval(() => this.fetchAgentStatuses(), this.pollingInterval * 1000);
  }

  rebuildEnv(): void {
    console.log('rebuilding env...');
  }
}
