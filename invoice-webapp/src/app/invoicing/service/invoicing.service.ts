import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry, map } from 'rxjs/operators';
import { Invoice } from '../invoice-form/invoice-form.component';
import { Location, DatePipe } from '@angular/common';


@Injectable({
  providedIn: 'root'
})
export class InvoicingService {
  invoice: Invoice
  constructor(private location: Location, private http: HttpClient) {
    this.invoice = new Invoice();
  }

  // Implements HTTP POST to create or update invoice that is displayed.
  addInvoice(invoice: Invoice) {
    console.log('addInvoice', invoice);
    this.http.post('/v1/invoice/save', invoice).subscribe(result => {
      console.log('this is the result: ', result.toString());
    })
    console.log('this is what were sending: ', JSON.stringify(invoice));
  }

  // Implements HTTP GET to fetch invoice by invoice number.
  getInvoice(invoiceNumber: number): Observable<Invoice> {
    invoiceNumber = 12345
    console.log('fetchInvoice', invoiceNumber);
    return this.http.get<Invoice>(`/v1/invoice/view/${invoiceNumber}`);
  }
}
