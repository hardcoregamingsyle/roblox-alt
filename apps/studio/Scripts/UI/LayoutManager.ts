export class LayoutManager {
  private readonly MIN_SIZE = { x: 200, y: 150 };
  private readonly MAX_SIZE = { x: 4096, y: 4096 };

  public Clamp(size: { x: number, y: number }): { x: number, y: number } {
    return {
      x: Math.max(this.MIN_SIZE.x, Math.min(this.MAX_SIZE.x, size.x)),
      y: Math.max(this.MIN_SIZE.y, Math.min(this.MAX_SIZE.y, size.y))
    };
  }
}