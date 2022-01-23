import { AccountsService } from './../accounts.service';
import { Account } from './../entities/account.entity';
import { Injectable, Scope } from '@nestjs/common';

/*Serviço no Node é normalmente compartilhado por aplicação igual no spring*/
/*Com esse Scope.REQUEST cada requisição feita terá uma instancia própria*/
@Injectable({ scope: Scope.REQUEST })
export class AccountStorageService {
  private _account: Account | null = null;

  constructor(private accountService: AccountsService) {}

  get account() {
    return this._account;
  }

  async setBy(token: string) {
    this._account = await this.accountService.findOne(token);
  }
}
