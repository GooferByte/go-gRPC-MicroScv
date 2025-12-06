# Schema & Validation Design

## Fields (purpose)
- `invoice_number`: Unique identifier assigned by the seller for traceability.
- `invoice_date`: Date the invoice is issued; anchors payment terms and timelines.
- `seller_name` / `buyer_name`: Parties to the transaction for billing and compliance.
- `currency`: Currency code (e.g., USD, EUR, GBP, INR) to interpret monetary fields.
- `line_items[]`: Collection of billed items; each has `description`, `quantity`, and `unit_price`.
- `net_total`: Sum of all line item totals before tax or fees.
- `tax_amount`: Total tax applied to the invoice.
- `gross_total`: Final amount due (net total plus tax).
- `due_date`: When payment is expected; derived from terms.

## Validation Rules (with rationale)
- Completeness: `invoice_number`, `invoice_date`, `seller_name`, `buyer_name`, and at least one `line_item` must be present to make the invoice actionable and auditable.
- Format: `invoice_date` and `due_date` must parse as valid dates within a sensible range (2000-01-01 to 2100-01-01) to avoid bad inputs and temporal anomalies.
- Format: `currency` must be in an allowed set (USD, EUR, GBP, INR, etc.) so downstream systems price correctly.
- Business: Sum of line item totals (`quantity * unit_price`) must equal `net_total` within a small tolerance to detect math or entry errors.
- Business: `gross_total` must equal `net_total + tax_amount` within tolerance to ensure totals roll up properly.
- Business: `due_date` must be on or after `invoice_date` so payment terms are non-negative.
- Anomaly/Duplicate: No duplicate invoice with the same `(invoice_number, seller_name, invoice_date)` to prevent double billing.
- Anomaly: Monetary totals (`net_total`, `gross_total`, `tax_amount`) must be non-negative and reasonable relative to line items to catch outliers or incorrect signs.
