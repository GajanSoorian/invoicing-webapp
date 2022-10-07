import { TestBed } from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http';

import { InvoicingService } from './invoicing.service';

describe('InvoicingService', () => {
  let service: InvoicingService;

  beforeEach(() => {
    TestBed.configureTestingModule({imports: [
      HttpClientModule]});
    service = TestBed.inject(InvoicingService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
