export type AircraftType = 'ATR' | 'Airbus 320' | 'Boeing 737 Max';

export interface VoucherFormState {
  name: string;
  id: string;
  flightNumber: string;
  date: string;
  aircraft: AircraftType;
}

export interface GenerateVoucherResponse {
  seats: string[];
}

export interface CheckVoucherResponse {
  exists: boolean;
}