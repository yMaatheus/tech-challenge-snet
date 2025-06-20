export interface Store {
  number: string;
  name: string;
  corporate_name: string;
  address: string;
  address_number: string;
  city: string;
  state: string;
  zip_code: string;
}

export interface StoreWithEstablishment extends Store {
  id: string;
  businessId: string;
  establishmentId: string;
}