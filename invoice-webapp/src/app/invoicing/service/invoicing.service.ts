import { Injectable, OnDestroy } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, Subscription } from 'rxjs';
import { Invoice } from '../invoice-form/invoice-form.component';


@Injectable({
  providedIn: 'root'
})
export class InvoicingService implements OnDestroy{

  invoice: Invoice;

  constructor(private _httpClient: HttpClient) {
    this.invoice = new Invoice();
  }
  ngOnDestroy(): void {
    throw new Error('Method not implemented.');
  }

  // Implements HTTP POST to create a new or update an existing invoice.
  addInvoice(invoice: Invoice) {
    this._httpClient.post('/v1/invoice/save', invoice).subscribe(result => {  })
  }

  // Implements HTTP GET to fetch invoice by invoice number.
  getInvoice(invoiceNumber: number): Observable<Invoice> {
    return this._httpClient.get<Invoice>(`/v1/invoice/view/${invoiceNumber}`);
  }
}
