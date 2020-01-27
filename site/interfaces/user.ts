export default interface User {
  id: string;
  createdAt: string;
  updatedAt: string;
  firstname: string;
  lastname: string;
  email: string;
  stripe_customer_id: string;
  type: string;
}
