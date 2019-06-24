import { Component, Input} from '@angular/core';

@Component({
  selector: 'rs-levelSensor',
  templateUrl: './levelSensor.component.html',
  styleUrls: []
})
export class LevelSensorComponent {
  @Input() sensor;
}
