import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CustomerSummary } from './customer-summary';

describe('CustomerSummary', () => {
  let component: CustomerSummary;
  let fixture: ComponentFixture<CustomerSummary>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CustomerSummary],
    }).compileComponents();

    fixture = TestBed.createComponent(CustomerSummary);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
