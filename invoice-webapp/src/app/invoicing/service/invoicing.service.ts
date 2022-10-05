import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError, retry, map } from 'rxjs/operators';
import {Invoice} from '../invoice-form/invoice-form.component';
import { Location } from '@angular/common';


@Injectable({
  providedIn: 'root'
})
export class InvoicingService {
  invoice:Invoice
  constructor(private location: Location, private http:HttpClient) {
    this.invoice = new Invoice();
   }

  simpleFunctionCall(): void {
    console.log('simpleFunctionCall');
  }


  addInvoice(invoice:Invoice){
    console.log('addInvoice', invoice);
   // this.http.post('/v1/invoice-create',JSON.stringify(invoice)).subscribe(result => {
   // })
   this.http.post('/v1/invoice-create',invoice).subscribe(result => {
    console.log('this is the result: ', result.toString());
     })
    console.log('this is what were sending: ', JSON.stringify(invoice));
  }

   getInvoice(invoiceNumber:number): Observable<Invoice>{
    invoiceNumber =12345
    console.log('fetchInvoice', invoiceNumber);

    /*const url = '/v1/invoice-display';
    return this.http.get<Invoice>(url).pipe(
      tap(_ => this.log(`fetched hero id=${id}`)),
      catchError(this.handleError<Hero>(`getHero id=${id}`))
    );
*/
    return this.http.get<Invoice>(`/v1/invoice-display/${invoiceNumber}`);


   }

/*
  loadInvoice(invoice:Invoice){
    this.invoice = this.http.get<Invoice[]>('/v1//invoice-display').subscribe
 }*/
}
