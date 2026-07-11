import React, { useState } from 'react';
// Menggunakan "import type" sesuai aturan TypeScript Vite terbaru
import type { VoucherFormState, AircraftType } from './types';
import { VoucherService } from './api';

export default function VoucherGenerator() {
  const [formData, setFormData] = useState<VoucherFormState>({
  name: '',
  id: '',
  flightNumber: '',
  date: '',
  aircraft: 'ATR',
});

  const [status, setStatus] = useState<'idle' | 'loading' | 'success' | 'error'>('idle');
  const [errorMessage, setErrorMessage] = useState<string>('');
  const [assignedSeats, setAssignedSeats] = useState<string[]>([]);

  // Mengubah format YYYY-MM-DD (dari input type="date") menjadi DD-MM-YYYY untuk API
  const formatDateToDDMMYYYY = (dateString: string) => {
    if (!dateString) return '';
    const [year, month, day] = dateString.split('-');
    return `${day}-${month}-${year}`;
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setStatus('loading');
    setErrorMessage('');
    
    const formattedDate = formatDateToDDMMYYYY(formData.date);

    try {
      // 1. Cek ketersediaan voucher
      const checkResult = await VoucherService.check(formData.flightNumber, formattedDate);

      if (checkResult.exists) {
        // 3. Jika sudah ada, tampilkan pesan error
        setStatus('error');
        setErrorMessage('Vouchers have already been generated for that flight and date.');
        return;
      }

      // 2. Jika belum ada, buat voucher baru
      const payload = { ...formData, flightDate: formattedDate };
      const generateResult = await VoucherService.generate(payload);

      setAssignedSeats(generateResult.seats);
      setStatus('success');
    } catch (error) {
      setStatus('error');
      setErrorMessage('An unexpected error occurred while processing your request.');
    }
  };

  // Tailwind class variables agar penulisan komponen JSX lebih rapi
  const inputClasses = "mt-1 block w-full rounded-md border border-gray-300 p-2 shadow-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500 disabled:bg-gray-100 disabled:text-gray-500";
  const labelClasses = "block text-sm font-medium text-gray-700";

  return (
    <div className="mx-auto mt-10 max-w-lg rounded-xl bg-white p-6 shadow-lg sm:p-8">
      <h1 className="mb-6 text-2xl font-bold text-gray-900">Generate Crew Vouchers</h1>

      <form onSubmit={handleSubmit} className="space-y-5">
        <div>
          <label className={labelClasses}>Crew Name</label>
          <input
            type="text"
            name="name"
            required
            value={formData.name}
            onChange={handleInputChange}
            disabled={status === 'loading'}
            className={inputClasses}
          />
        </div>

        <div>
          <label className={labelClasses}>Crew ID</label>
          <input
            type="text"
            name="id"
            required
            value={formData.id}
            onChange={handleInputChange}
            disabled={status === 'loading'}
            className={inputClasses}
          />
        </div>

        <div>
          <label className={labelClasses}>Flight Number</label>
          <input
            type="text"
            name="flightNumber"
            required
            placeholder="e.g., GA102"
            value={formData.flightNumber}
            onChange={handleInputChange}
            disabled={status === 'loading'}
            className={inputClasses}
          />
        </div>

        <div>
          <label className={labelClasses}>Flight Date</label>
          <input
            type="date"
            name="date"
            required
            value={formData.date}
            onChange={handleInputChange}
            disabled={status === 'loading'}
            className={inputClasses}
          />
        </div>

        <div>
          <label className={labelClasses}>Aircraft Type</label>
          <select
            name="aircraft"
            value={formData.aircraft}
            onChange={handleInputChange}
            disabled={status === 'loading'}
            className={inputClasses}
          >
            <option value="ATR">ATR</option>
            <option value="Airbus 320">Airbus 320</option>
            <option value="Boeing 737 Max">Boeing 737 Max</option>
          </select>
        </div>

        <button
          type="submit"
          disabled={status === 'loading'}
          className="mt-4 w-full rounded-md bg-blue-600 px-4 py-3 text-sm font-semibold text-white shadow-sm transition-colors hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 disabled:cursor-not-allowed disabled:bg-blue-400"
        >
          {status === 'loading' ? 'Processing...' : 'Generate Vouchers'}
        </button>
      </form>

      {/* Area Pesan Error */}
      {status === 'error' && (
        <div className="mt-6 rounded-md border border-red-200 bg-red-50 p-4">
          <p className="text-sm font-medium text-red-800">{errorMessage}</p>
        </div>
      )}

      {/* Area Pesan Sukses */}
      {status === 'success' && (
        <div className="mt-6 rounded-md border border-green-200 bg-green-50 p-4">
          <h3 className="text-sm font-medium text-green-800">Vouchers Generated Successfully!</h3>
          <div className="mt-2 text-sm text-green-700">
            <p className="mb-2">Assigned Seat Numbers:</p>
            <div className="flex flex-wrap gap-2">
              {assignedSeats.map((seat, index) => (
                <span 
                  key={index} 
                  className="inline-flex items-center rounded-md bg-green-200 px-2.5 py-0.5 text-xs font-bold text-green-900"
                >
                  {seat}
                </span>
              ))}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}