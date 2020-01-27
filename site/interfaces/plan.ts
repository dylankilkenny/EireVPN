export default interface Plan {
  id: string;
  createdAt: string;
  updatedAt: string;
  name: string;
  amount: number;
  interval: string;
  interval_count: number;
  plan_type: string;
  currency: string;
  stripe_plan_id: string;
  stripe_product_id: string;
}
