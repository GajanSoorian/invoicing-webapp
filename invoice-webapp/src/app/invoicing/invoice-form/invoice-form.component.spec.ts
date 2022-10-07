import { FormsModule } from '@angular/forms';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { InvoicingService } from './../service/invoicing.service';
import { HttpClientModule } from '@angular/common/http';
import { DebugElement } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';

import { InvoiceFormComponent } from './invoice-form.component';

describe('InvoiceFormComponent', () => {
  let component: InvoiceFormComponent;
  let fixture: ComponentFixture<InvoiceFormComponent>;
  let invoiceService: InvoicingService;
  let createButtonDebugElement: DebugElement;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [InvoiceFormComponent],
      imports: [
        HttpClientModule,
        MatSnackBarModule,
        FormsModule
      ]
    })
      .compileComponents();

    fixture = TestBed.createComponent(InvoiceFormComponent);
    invoiceService = TestBed.inject(InvoicingService);
    component = fixture.componentInstance;
    fixture.detectChanges();

    createButtonDebugElement = fixture.debugElement.query(By.css('.create-button'));
  });

  fit('should create InvoiceFormComponent', () => {
    expect(component).toBeTruthy();
  });

  fit('should create new empty invoice object when create is clicked', () => {
    createButtonDebugElement.triggerEventHandler('click');
    expect(component.invoice).toBeTruthy();
    expect(component.invoice.products.length).toEqual(1);
    expect(component.invoice.products[0].price).toEqual(0);
    expect(component.invoice.products[0].description).toEqual('');
  });

});
