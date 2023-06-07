assets = $(wildcard *.asset.tape) $(wildcard **/*.asset.tape)
gifs   = $(assets:.asset.tape=.gif)

prefix_file = prefix.tape
suffix_file = suffix.tape

all: $(gifs)

%.gif: %.asset.tape
	@cat ${prefix_file} >  temp.tape
	@cat $<             >> temp.tape
	@cat ${suffix_file} >> temp.tape
	@echo vhs $<
	@vhs temp.tape -o $@

clean:
	-rm *.gif
	-rm temp.tape

.PHONY: all
