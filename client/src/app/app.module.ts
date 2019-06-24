import { BrowserModule } from '@angular/platform-browser';
import { NgModule, NO_ERRORS_SCHEMA  } from '@angular/core';

import { AppComponent } from './app.component';
import { InfoComponent } from './info.component';
import {MaintenanceComponent} from "./maintenance.component";
import {ProbeComponent} from "./probe.component";
import {LevelSensorComponent} from "./levelSensor.component";
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent,
    InfoComponent,
    MaintenanceComponent,
    ProbeComponent,
    LevelSensorComponent,
  ],
  imports: [
    BrowserModule,
    NgbModule.forRoot(),
    HttpClientModule,
  ],
  schemas: [ NO_ERRORS_SCHEMA ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
