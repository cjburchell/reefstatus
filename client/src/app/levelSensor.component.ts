import { Component, Input} from '@angular/core';
import {LevelSensor} from './controller.service';

@Component({
  selector: 'app-level-sensor',
  templateUrl: './levelSensor.component.html',
  styleUrls: []
})
export class LevelSensorComponent {
  @Input() sensor: LevelSensor;

  startWaterChange() {

  }

  clearAlarm() {
  }
}
