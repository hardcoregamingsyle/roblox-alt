import { PrismaClient } from '@prisma/client';

export const applyTenantMiddleware = (client: PrismaClient, tenantId: string) => {
  client.$use(async (params, next) => {
    const protectedModels = ['Game', 'User', 'Asset'];
    
    if (protectedModels.includes(params.model!)) {
      // FIX: Enforce tenant_id on ALL CRUD operations, preventing bypass
      params.args = params.args || {};
      params.args.where = {
        ...params.args.where,
        tenant_id: tenantId,
      };
      
      // Prevent manual override of tenant_id in update/create operations
      if (params.args.data) {
        params.args.data.tenant_id = tenantId;
      }
    }
    return next(params);
  });
};