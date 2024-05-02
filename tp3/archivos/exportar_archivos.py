LATITUD = 0
LONGITUD = 1

def exportar_kml(ruta, camino, coordenadas):
    ap_anterior = ''
    
    with open(ruta, 'w') as f:
        f.write('<?xml version="1.0" encoding="UTF-8"?>\n')
        f.write('<kml xmlns="http://earth.google.com/kml/2.1">\n')
        f.write('\t<Document>\n')
        f.write('\t\t<name>Recorrido del viaje a realizar</name>\n')
        for ap in camino:
            f.write(f'\t\t<Placemark>\n')
            f.write(f'\t\t\t<name>{ap}</name>\n')
            f.write(f'\t\t\t<Point>\n')
            f.write(f'\t\t\t\t<coordinates>{coordenadas[ap][LONGITUD]}, {coordenadas[ap][LATITUD]}</coordinates>\n')
            f.write(f'\t\t\t</Point>\n')
            f.write(f'\t\t</Placemark>\n')
            ap_anterior = ap    
                
        for p, ap in enumerate(camino):
            if p == 0:
                ap_anterior = ap
            else:
                f.write(f'\t\t<Placemark>\n')
                f.write(f'\t\t\t<LineString>\n')
                f.write(f'\t\t\t\t<coordinates>{coordenadas[ap_anterior][LONGITUD]}, {coordenadas[ap_anterior][LATITUD]} {coordenadas[ap][LONGITUD]}, {coordenadas[ap][LATITUD]}</coordinates>\n')
                f.write(f'\t\t\t</LineString>\n')
                f.write(f'\t\t</Placemark>\n')
                ap_anterior = ap
                
        f.write('\t</Document>\n')
        f.write('</kml>\n')

def exportar_nueva_aerolinea(ruta, rutas_minimas):
    with open(ruta, 'w') as f:
        for origen, destino, pesos in rutas_minimas:
            f.write(f'{origen},{destino},{pesos[0]},{pesos[1]},{pesos[2]}\n')
