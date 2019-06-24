import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'rs-maintenance',
  templateUrl: './maintenance.component.html',
  styleUrls: []
})
export class MaintenanceComponent {
  @Input() maintenance;
  @Output() triggered = new EventEmitter();

  onMaintenance(){
    this.triggered.emit(this.maintenance);
    console.log("Maintenance! " + this.maintenance.DisplayName)
  }
}
