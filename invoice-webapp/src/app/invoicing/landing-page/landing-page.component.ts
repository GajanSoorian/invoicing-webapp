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

  onCreateClick(): void {
    this._router.navigateByUrl('/create');
  }
  onViewClick(): void {

  }
}
