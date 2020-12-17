import {
    Directive,
    Input,
    OnInit,
    TemplateRef,
    ViewContainerRef
  } from '@angular/core';
  
  import { UserService } from '../services/user.service';
  
  @Directive({ selector: '[appShowAdmin]' })
  export class ShowAdminDirective implements OnInit {
    constructor(
      private templateRef: TemplateRef<any>,
      private userService: UserService,
      private viewContainer: ViewContainerRef
    ) {}
  
    condition: boolean;
  
    ngOnInit() {
      let isAdmin = this.userService.isAdmin();

      if(isAdmin && this.condition || !isAdmin && !this.condition){
          this.viewContainer.createEmbeddedView(this.templateRef);
      }else{
        this.viewContainer.clear();
      }
    }
  
    @Input() set appShowAdmin(condition: boolean) {
      this.condition = condition;
    }
  
  }
  