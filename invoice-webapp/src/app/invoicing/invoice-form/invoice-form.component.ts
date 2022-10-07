import { Observable, Subscription } from 'rxjs';
import { Component } from '@angular/core';
import { InvoicingService } from '../service/invoicing.service';
import { MatSnackBar } from '@angular/material/snack-bar';
import * as _ from 'lodash';

@Component({
  selector: 'app-invoice-form',
  templateUrl: './invoice-form.component.html',
  styleUrls: ['./invoice-form.component.less']
})
export class InvoiceFormComponent {

  customerName: string = ''
  email: string = ''
  name: string = ''
  description: string = ''
  price: number = 0
  dueDate: Date = new Date()
  isViewModifyMode = false
  isCreateMode = false;
  invoiceNumberToFetch: number = 0

  invoice = new Invoice();

  constructor(
    private _invoicingService: InvoicingService,
    private _snackBar: MatSnackBar
  ) { }

  // Add line items to the invoice using the '+' button.
  addProduct() {
    this.invoice.products.push({
      name: this.name,
      description: this.description,
      price: this.price
    }) ;
  }

  // Remove line items from the invoice using the '-' button.
  removeProduct() {
    this.invoice.products.pop();
    this.invoice.totalAmount = (this.invoice.products.reduce((n, {price}) => n + price!, 0))
  }

  // Send request to the service to create or update an invoice.
  onSubmitClick(): void {
    let message = '';
    let messageType ='info';
    if(this.canSubmitInvoice()) {
      this._invoicingService.addInvoice(this.invoice)
      message = `Invoice with number ${this.invoice.invoiceNumber} created successfully`
    } else {
      message = `Please add required details and products to create an invoice`
      messageType = 'warn'
    }

    this.popupMessage(message, "Ok" , messageType)
  }

  private canSubmitInvoice(): boolean {
    return !_.isNil(this.invoice.customerName)
    && this.invoice.customerName!.length > 0
    && !_.isNil(this.invoice.email)
    && this.invoice.email!.length > 0
    && this.invoice.products.length > 0
    && !_.isNil(this.invoice.products[0].name)
    && this.invoice.products[0].name!.length > 0
  }

  onCreateClick(): void {
    this.invoice = new Invoice();
    this.isCreateMode =true
    this.isViewModifyMode =false

    if( this.invoice.products.length == 0){
    this.invoice.products.push({
      name:"",
      description: "",
      price: 0})
    this.invoice.totalAmount =0
    }
    this.invoice.invoiceNumber = (new Date).getTime()
  }

  onProductPriceUpdate(): void {
    this.invoice.totalAmount = (this.invoice.products.reduce((n, {price}) => n + price!, 0))
  }

  onViewModifyClick(): void {
    this.isViewModifyMode =true
    this.isCreateMode =false
    this.invoice = new Invoice();
  }

  onSearchClick(): void {
    this.isCreateMode =true
    this.isViewModifyMode =false
    this._invoicingService.getInvoice(this.invoiceNumberToFetch).subscribe (data => {
      this.invoice = data;
    })
  }

  popupMessage(message: string, action: string, msgType: string) {
    this._snackBar.open(message, action, { panelClass: msgType });
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
