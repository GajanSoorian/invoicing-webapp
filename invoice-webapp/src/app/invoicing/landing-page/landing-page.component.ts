import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.less']
})
export class LandingPageComponent implements OnInit {

  constructor(private _router: Router) { }

  ngOnInit(): void {
  }

  // Can be used to route to pages for create action. Not used at this time.
  onCreateClick(): void {
    this._router.navigateByUrl('/create');
  }
   // Can be used to route to pages for view or modify actions. Not used at this time.
  onViewClick(): void {

  }
}
