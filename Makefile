MAPS = mch1.png mch2.png

.PHONY: maps
maps: $(MAPS)

%.png: %.dot
	dot -Tpng $< > $@

clean:
	rm -f $(MAPS)
