#!/usr/bin/python3
import json

with open("output.sql", "w") as sqlOut:
    for p in ["woka.json", "companions.json"]:
        with open(p, "r") as textures:
            data = json.load(textures)

            for layer in data.keys():
                for layerCollection in data[layer]["collections"]:
                    for layerTexture in layerCollection["textures"]:
                        sqlOut.write(f"INSERT INTO textures (texture, layer, url, tags, created_at) VALUES('{layerTexture['id']}', '{layer}', '%FRONTEND_URL%/public/{layerTexture['url']}', '[]', NOW());\n")