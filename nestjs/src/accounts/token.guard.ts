import { CanActivate, ExecutionContext, Injectable } from '@nestjs/common';
import { AccountStorageService } from './account-storage/account-storage.service';

@Injectable()
export class TokenGuard implements CanActivate {
  constructor(private accountStorage: AccountStorageService) { }

  async canActivate(context: ExecutionContext): Promise<boolean> {
    if (context.getType() !== 'http') {
      console.log('Gardiao Token liberadao para o tipo: ' + context.getType());
      return true;
    }
    const request = context.switchToHttp().getRequest();
    const token = request.headers?.['x-token'] as string;
    try {
      if (token) {
        await this.accountStorage.setBy(token);
        return true;
      }
    } catch (error) {
      console.error(error);
      return false;
    }

    return false;
  }
}
