export interface Server {
  id: string;
  name: string;
  description: string;
  fqdn: string;
  ip_address: string;
  orgs_count?: string;
}
