import { Component } from '@angular/core';
import {ControllerService, Info, LevelSensor, Probe} from './controller.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  constructor(private controller: ControllerService) {
    this.refresh();
  }

  public isCollapsed = false;
  public info: Info;

  public probes: Probe[] = [];
  public levelSensors: LevelSensor[] = [];
  isCollapsedLevelSensors: boolean;

  refresh() {
    this.controller.getInfo().subscribe((info: Info) => this.info = info);
    this.controller.getProbes().subscribe((probes: Probe[]) => this.probes = probes);
    this.controller.getLevelSesnsors().subscribe((levelSensors: LevelSensor[]) => this.levelSensors = levelSensors);
  }
}
