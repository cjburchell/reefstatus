import { Component, Input, Output, EventEmitter } from '@angular/core';
import {Maintenance} from './controller.service';

@Component({
  selector: 'app-maintenance',
  templateUrl: './maintenance.component.html',
  styleUrls: []
})
export class MaintenanceComponent {
  @Input() maintenance: Maintenance;
  @Output() triggered = new EventEmitter();

  onMaintenance() {
    this.triggered.emit(this.maintenance);
    console.log('Maintenance! ' + this.maintenance.DisplayName);
  }
}
