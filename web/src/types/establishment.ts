import type { StoreWithEstablishment } from "~/types/store"

export interface Establishment {
  id: string
  number: string
  name: string
  corporate_name: string
  address: string
  address_number: string
  city: string
  state: string
  zip_code: string
}

export interface EstablishmentWithStores extends Establishment {
  stores: StoreWithEstablishment[]
}
