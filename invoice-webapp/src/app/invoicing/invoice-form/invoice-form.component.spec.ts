import { DebugElement } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';

import { InvoiceFormComponent } from './invoice-form.component';

describe('InvoiceFormComponent', () => {
  let component: InvoiceFormComponent;
  let fixture: ComponentFixture<InvoiceFormComponent>;
  let createButtonDebugElement: DebugElement;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InvoiceFormComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InvoiceFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();

    createButtonDebugElement = fixture.debugElement.query(By.css('create-button'));
  });

  it('should create InvoiceFormComponent', () => {
    expect(component).toBeTruthy();
  });

  it('should create new empty invoice object when create is clicked', () => {
    expect(component.invoice).toBeFalsy();
    createButtonDebugElement.triggerEventHandler('click');
    expect(component.invoice).toBeTruthy();
  });

});
