import { Component, Input} from '@angular/core';
import {Probe} from '../services/contracts/contracts';


@Component({
  selector: 'app-probe',
  templateUrl: './probe.component.html',
  styleUrls: []
})
export class ProbeComponent {
  @Input() probe: Probe | undefined;
}
