import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AccountSelectionComponent } from './account-selection';

describe('AccountSelection', () => {
  let component: AccountSelectionComponent;
  let fixture: ComponentFixture<AccountSelectionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AccountSelectionComponent],
    }).compileComponents();

    fixture = TestBed.createComponent(AccountSelectionComponent);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
