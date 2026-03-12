import { api } from "./api";
import type { APIResponse } from "../types/auth";
import type { Product, CreateProductRequest } from "../types/product";

export const productService = {
  getAll: async (): Promise<Product[]> => {
    const { data } = await api.get<APIResponse<Product[]>>("/api/v1/products");
    return data.data ?? [];
  },

  getById: async (id: number): Promise<Product> => {
    const { data } = await api.get<APIResponse<Product>>(`/api/v1/products/${id}`);
    return data.data;
  },

  create: async (req: CreateProductRequest): Promise<Product> => {
    const { data } = await api.post<APIResponse<Product>>("/api/v1/products", req);
    return data.data;
  },

  update: async (id: number, req: Partial<CreateProductRequest>): Promise<Product> => {
    const { data } = await api.put<APIResponse<Product>>(`/api/v1/products/${id}`, req);
    return data.data;
  },

  delete: async (id: number): Promise<void> => {
    await api.delete(`/api/v1/products/${id}`);
  },
};
