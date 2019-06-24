import { Component } from '@angular/core';
import {ControllerService, Info} from "./controller.service";

@Component({
  selector: 'rs-app',
  templateUrl: './app.component.html',
  styleUrls: []
})
export class AppComponent {

  constructor(private controller: ControllerService) {
    this.refresh();
  }

  refresh(){
    this.controller.getInfo().subscribe((info: Info) => this.info = info);
  }

  isCollapsed = false;
  info : Info;

  probes = [{
    DisplayName : "Temp",
    AlarmState: 'On',
    ConvertedValue: 22,
    Units: "C"
  }, {
    DisplayName : "PH",
    AlarmState: 'Off',
    ConvertedValue: 12,
    Units: "PH"

  }];
  levelSensors= [];
  sports = [];
  lports = [];
}
