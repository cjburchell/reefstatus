import { Component, Input, Output, EventEmitter } from '@angular/core';
import {Info} from '../services/contracts/contracts';

@Component({
  selector: 'app-info',
  templateUrl: './info.component.html',
  styleUrls: []
})
export class InfoComponent {
 @Input() info: Info | undefined;
 @Output() feedPause = new EventEmitter();
 @Output() thunderstorm = new EventEmitter();

  onFeedPause(): void {
    this.feedPause.emit();
    console.log('FeedPause!');
  }

  onThunderstorm(): void {
    this.thunderstorm.emit();
    console.log('Thunderstorm!');
  }
}
