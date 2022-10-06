import { Component, OnInit, NgModule } from '@angular/core';
import { InvoicingService } from '../service/invoicing.service';
import { MatSnackBar } from '@angular/material/snack-bar';
import {FormGroup} from '@angular/forms';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';

@Component({
  selector: 'app-invoice-form',
  templateUrl: './invoice-form.component.html',
  styleUrls: ['./invoice-form.component.less']
})
export class InvoiceFormComponent implements OnInit {

  customerName: string = ''
  email: string = ''
  name: string = ''
  description: string = ''
  price: number = 0
  dueDate: Date = new Date()
  isViewModifyMode = false
  isCreateMode = false;
  invoiceNumberToFetch: number =0

  invoice = new Invoice();
  constructor(
    private _invoicingService: InvoicingService,
    private _snackBar: MatSnackBar
  ) { }

  ngOnInit(): void {
  }

  // Add line items to the invoice using the '+' button.
  addProduct() {
    this.invoice.products.push({
      name: this.name,
      description: this.description,
      price: this.price
    }) ;
    this.invoice.totalAmount = (this.invoice.products.reduce((n, {price}) => n + price!, 0))
  }

  // Remove line items from the invoice with the '-' button.
  removeProduct() {
    this.invoice.products.pop();
    this.invoice.totalAmount = (this.invoice.products.reduce((n, {price}) => n + price!, 0))
  }

  onSubmitClick(): void {
    this.invoice.customerName = this.customerName
    this.invoice.email = this.email
    this.invoice.dueByDate = this.dueDate
    this.invoice.invoiceNumber
    console.log("Need to sed this to the back end : ",this.invoice.dueByDate ,  this.invoice.customerName,
    this.invoice.email, this.invoice.invoiceNumber, this.invoice.products);
    this._invoicingService.addInvoice(this.invoice)
    let message = `Invoice with number ${this.invoice.invoiceNumber} created successfully`
    this.popupMessage(message, "Ok")
    //Todo : reset the form and component.
  }

  onCreateClick(): void {
    this.invoice.invoiceNumber = (new Date).getTime()
    this.isCreateMode =true
    this.isViewModifyMode =false
    if( this.invoice.products.length == 0){
    this.invoice.products.push({
      name:"",
      description: "",
      price: 0})
    this.invoice.totalAmount =0
    }
  }
  onViewClick(): void {
    this.isViewModifyMode =true
    this.invoice.invoiceNumber = 0
}

  onSearchClick(): void {
    this.invoiceNumberToFetch = this.invoice.invoiceNumber!
    console.log("Invoice to fetch", this.invoiceNumberToFetch)
    this._invoicingService.getInvoice(this.invoiceNumberToFetch).subscribe (data => {
      console.log(data);
      this.invoice = data;
    })
  }

  popupMessage(message: string, action: string) {
    this._snackBar.open(message, action);
  }


  handleError(err: any) {
    this._snackBar.open(`Action Failed.`, 'Ok', { panelClass: 'warn' });
    return new Error(err);
  }

}
interface Item {
  name: string | undefined
  description: string | undefined
  price: number | undefined
}

export class Invoice {
  invoiceNumber : number | undefined
  customerName: string | undefined
  email: string | undefined

  products: Item[] = []
  dueByDate: Date | undefined
  totalAmount: number | undefined

  constructor() {
  }
}
