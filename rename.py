import glob
import os
from tqdm import tqdm

for filename in tqdm(glob.glob("blast-geth/**", recursive=True)):
    if os.path.isdir(filename):
        continue

    with open(filename, "rb") as f:
        content = f.read()

    newContent = content.replace(
        b"github.com/ethereum/go-ethereum",
        b"github.com/tenderly/blast/blast-geth",
    )

    if newContent != content:
        with open(filename, "wb") as f:
            f.write(newContent)
