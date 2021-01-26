import { Component } from '@angular/core';
import {ControllerService} from './services/controller.service';
import {Info, LevelSensor, Probe} from './services/contracts/contracts';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  constructor(private controller: ControllerService) {
    this.refresh().then(() => {});
  }

  public isCollapsed = false;
  public info: Info | undefined;

  public probes: Probe[] = [];
  public levelSensors: LevelSensor[] = [];
  isCollapsedLevelSensors: boolean | undefined;

  async refresh(): Promise<void> {
    this.info = await this.controller.getInfo();
    this.probes = await this.controller.getProbes();
    this.levelSensors = await this.controller.getLevelSesnsors();
  }
}
