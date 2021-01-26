import { Component, Input, Output, EventEmitter } from '@angular/core';
import {Maintenance} from '../services/contracts/contracts';

@Component({
  selector: 'app-maintenance',
  templateUrl: './maintenance.component.html',
  styleUrls: []
})
export class MaintenanceComponent {
  @Input() maintenance: Maintenance | undefined;
  @Output() triggered = new EventEmitter();

  onMaintenance(): void {
    if (this.maintenance){
      this.triggered.emit(this.maintenance);
      console.log('Maintenance! ' + this.maintenance.DisplayName);
    }
  }
}
