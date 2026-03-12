import { useEffect, useState, type FormEvent } from "react";
import { Plus, Trash2 } from "lucide-react";
import { productService } from "../../services/productService";
import type { Product, CreateProductRequest } from "../../types/product";
import styles from "./Product.module.css";

const emptyForm: CreateProductRequest = {
  name: "",
  description: "",
  price: 0,
  stock: 0,
};

export default function ProductPage() {
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);
  const [form, setForm] = useState<CreateProductRequest>(emptyForm);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchProducts = () => {
    setLoading(true);
    productService
      .getAll()
      .then((list) => setProducts(list ?? []))
      .finally(() => setLoading(false));
  };

  useEffect(() => {
    fetchProducts();
  }, []);

  const handleDelete = async (id: number) => {
    if (!window.confirm("Delete this product?")) return;
    await productService.delete(id);
    setProducts((prev) => prev.filter((p) => p.id !== id));
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setError(null);
    setSubmitting(true);
    try {
      const created = await productService.create(form);
      setProducts((prev) => [...prev, created]);
      setShowModal(false);
      setForm(emptyForm);
    } catch {
      setError("Failed to create product.");
    } finally {
      setSubmitting(false);
    }
  };

  const set = (field: keyof CreateProductRequest) =>
    (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) =>
      setForm((prev) => ({
        ...prev,
        [field]: field === "price" || field === "stock" ? Number(e.target.value) : e.target.value,
      }));

  return (
    <div className={styles.page}>
      <div className={styles.header}>
        <h1>Products</h1>
        <button className={styles.btnAdd} onClick={() => setShowModal(true)}>
          <Plus size={16} /> Add Product
        </button>
      </div>

      <div className={styles.tableWrapper}>
        <table className={styles.table}>
          <thead>
            <tr>
              <th>#</th>
              <th>Name</th>
              <th>Description</th>
              <th>Price</th>
              <th>Stock</th>
              <th>Created</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {loading ? (
              <tr>
                <td colSpan={7} className={styles.empty}>LoadingÅc</td>
              </tr>
            ) : products.length === 0 ? (
              <tr>
                <td colSpan={7} className={styles.empty}>No products yet. Add one!</td>
              </tr>
            ) : (
              products.map((p) => (
                <tr key={p.id}>
                  <td>{p.id}</td>
                  <td>{p.name}</td>
                  <td>{p.description || "?"}</td>
                  <td>${p.price.toFixed(2)}</td>
                  <td>
                    <span className={`${styles.badge} ${p.stock === 0 ? styles.badgeLow : ""}`}>
                      {p.stock}
                    </span>
                  </td>
                  <td>{new Date(p.created_at).toLocaleDateString()}</td>
                  <td>
                    <div className={styles.actions}>
                      <button
                        className={styles.btnDelete}
                        onClick={() => handleDelete(p.id)}
                      >
                        <Trash2 size={14} />
                      </button>
                    </div>
                  </td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>

      {showModal && (
        <div className={styles.overlay} onClick={() => setShowModal(false)}>
          <div className={styles.modal} onClick={(e) => e.stopPropagation()}>
            <h2>Add Product</h2>
            <form className={styles.form} onSubmit={handleSubmit}>
              <div className={styles.field}>
                <label htmlFor="name">Name *</label>
                <input
                  id="name"
                  value={form.name}
                  onChange={set("name")}
                  required
                  placeholder="Product name"
                />
              </div>
              <div className={styles.field}>
                <label htmlFor="description">Description</label>
                <textarea
                  id="description"
                  value={form.description}
                  onChange={set("description")}
                  rows={3}
                  placeholder="Optional description"
                />
              </div>
              <div className={styles.row}>
                <div className={styles.field}>
                  <label htmlFor="price">Price *</label>
                  <input
                    id="price"
                    type="number"
                    min="0"
                    step="0.01"
                    value={form.price}
                    onChange={set("price")}
                    required
                  />
                </div>
                <div className={styles.field}>
                  <label htmlFor="stock">Stock</label>
                  <input
                    id="stock"
                    type="number"
                    min="0"
                    value={form.stock}
                    onChange={set("stock")}
                  />
                </div>
              </div>
              {error && <p className={styles.error}>{error}</p>}
              <div className={styles.modalActions}>
                <button
                  type="button"
                  className={styles.btnCancel}
                  onClick={() => { setShowModal(false); setForm(emptyForm); setError(null); }}
                >
                  Cancel
                </button>
                <button type="submit" className={styles.btnSubmit} disabled={submitting}>
                  {submitting ? "SavingÅc" : "Save"}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}