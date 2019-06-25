import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import {InfoComponent} from './info.component';
import {LevelSensorComponent} from './levelSensor.component';
import {MaintenanceComponent} from './maintenance.component';
import {ProbeComponent} from './probe.component';
import {HttpClientModule} from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent,
    InfoComponent,
    LevelSensorComponent,
    MaintenanceComponent,
    ProbeComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
