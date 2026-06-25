import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BeneficiariesComponent } from './beneficiaries';

describe('Beneficiaries', () => {
  let component: BeneficiariesComponent;
  let fixture: ComponentFixture<BeneficiariesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BeneficiariesComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(BeneficiariesComponent);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
