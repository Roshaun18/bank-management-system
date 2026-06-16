import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CustomerSummaryComponent } from './customer-summary';

describe('CustomerSummary', () => {
  let component: CustomerSummaryComponent;
  let fixture: ComponentFixture<CustomerSummaryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CustomerSummaryComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(CustomerSummaryComponent);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
