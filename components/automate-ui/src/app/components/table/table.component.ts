import { Component, HostBinding } from '@angular/core';

@Component({
  selector: 'chef-table-new',
  templateUrl: './table.component.html',
  styleUrls: ['./table.component.scss']
})
export class TableComponent {
  @HostBinding('attr.role') role = 'table';
}
