import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'rs-info',
  templateUrl: './info.component.html',
  styleUrls: []
})
export class InfoComponent {
 @Input() info;
 @Output() feedPause = new EventEmitter();
 @Output() thunderstorm = new EventEmitter();

  onFeedPause(){
    this.feedPause.emit();
    console.log("FeedPause!")
  }

  onThunderstorm(){
    this.thunderstorm.emit();
    console.log("Thunderstorm!")
  }
}
