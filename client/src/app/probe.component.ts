import { Component, Input} from '@angular/core';
import {Probe} from './controller.service';

@Component({
  selector: 'app-probe',
  templateUrl: './probe.component.html',
  styleUrls: []
})
export class ProbeComponent {
  @Input() probe: Probe;
}
