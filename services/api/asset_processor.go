func HandleUpload(c *gin.Context) {
    // FIX: Limit request body size to prevent OOM
    c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 1024 * 1024)
    
    // ... existing logic ...
    // FIX: Ensure parameters are used in Exec
    _, err = tx.Exec(c.Request.Context(), "UPDATE assets SET status=$1, metadata=$2 WHERE id=$3", "ACTIVE", dto.Metadata, id)
}