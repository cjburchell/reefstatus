import { Component, Input, Output, EventEmitter } from '@angular/core';
import {Info} from './controller.service';

@Component({
  selector: 'app-info',
  templateUrl: './info.component.html',
  styleUrls: []
})
export class InfoComponent {
 @Input() info: Info;
 @Output() feedPause = new EventEmitter();
 @Output() thunderstorm = new EventEmitter();

  onFeedPause() {
    this.feedPause.emit();
    console.log('FeedPause!');
  }

  onThunderstorm() {
    this.thunderstorm.emit();
    console.log('Thunderstorm!');
  }
}
