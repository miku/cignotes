map.png: map.dot
	dot -Tpng $< > $@

clean:
	rm -f map.png

