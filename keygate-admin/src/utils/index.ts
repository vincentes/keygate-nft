export async function calculateEtag(blob: Blob): Promise<string> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onloadend = async () => {
      if (reader.result) {
        try {
          const arrayBuffer = reader.result as ArrayBuffer;
          const hashBuffer = await crypto.subtle.digest("SHA-256", arrayBuffer);
          const hashArray = Array.from(new Uint8Array(hashBuffer));
          const hashHex = hashArray
            .map((b) => b.toString(16).padStart(2, "0"))
            .join("");
          resolve(hashHex);
        } catch (error) {
          reject(error);
        }
      } else {
        reject(new Error("Failed to read blob"));
      }
    };
    reader.onerror = () => {
      reject(reader.error);
    };
    reader.readAsArrayBuffer(blob);
  });
}
