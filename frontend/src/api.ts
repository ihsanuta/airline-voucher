import type { VoucherFormState, CheckVoucherResponse, GenerateVoucherResponse } from './types';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || '/api';

export const VoucherService = {
  check: async (flightNumber: string, date: string): Promise<CheckVoucherResponse> => {
    const response = await fetch(`${API_BASE_URL}/check`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ flightNumber, date }),
    });
    if (!response.ok) throw new Error('Failed to check flight details');
    return response.json();
  },

  generate: async (data: VoucherFormState): Promise<GenerateVoucherResponse> => {
    const response = await fetch(`${API_BASE_URL}/generate`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });
    if (!response.ok) throw new Error('Failed to generate vouchers');
    return response.json();
  }
};