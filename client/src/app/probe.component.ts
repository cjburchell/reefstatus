import { Component, Input} from '@angular/core';

@Component({
  selector: 'rs-probe',
  templateUrl: './probe.component.html',
  styleUrls: []
})
export class ProbeComponent {
  @Input() probe;
}
