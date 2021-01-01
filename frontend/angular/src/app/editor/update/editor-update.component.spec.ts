import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditorUpdateComponent } from './editor-update.component';

describe('EditorUpdateComponent', () => {
  let component: EditorUpdateComponent;
  let fixture: ComponentFixture<EditorUpdateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EditorUpdateComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EditorUpdateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
