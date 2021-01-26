import { Component, Input} from '@angular/core';
import {LevelSensor} from '../services/contracts/contracts';


@Component({
  selector: 'app-level-sensor',
  templateUrl: './levelSensor.component.html',
  styleUrls: []
})
export class LevelSensorComponent {
  @Input() sensor: LevelSensor | undefined;

  startWaterChange(): void {

  }

  clearAlarm(): void {
  }
}
