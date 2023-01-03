package ini

import (
	"bufio"
	"strings"
)

type Format struct {
	// used for generating
	char_section_start string
	char_section_end   string
	char_assign        string
	char_comment       string
}

// used for parsing
func (f Format) is_section_start(ch string) bool { return ch == f.char_section_start }
func (f Format) is_section_end(ch string) bool   { return ch == f.char_section_end }
func (f Format) is_assign(ch string) bool        { return ch == f.char_assign }
func (f Format) is_comment(ch string) bool       { return ch == f.char_comment }

func newFormat(section_start string, section_end string, assign string, comment string) Format {
	return Format{
		char_section_start: section_start,
		char_section_end:   section_end,
		char_assign:        assign,
		char_comment:       comment,
	}
}

func NewFormat() Format {
	return newFormat("[", "]", "=", ";")
}

type Tuple_v struct {
	name  string
	value string
}

type Section struct {
	name string
	vars []Tuple_v
}

func (s Section) has(key string) bool {
	for _, v := range s.vars {
		if v.name == key {
			return true
		}
	}
	return false
}

func (s Section) at(key string) Tuple_v {
	for _, v := range s.vars {
		if v.name == key {
			return v
		}
	}
	return Tuple_v{"", ""}
}

/* if variable doesnt exist, add variable
 * else exit
 */
func (s *Section) insert(v Tuple_v) {
	if s.has(v.name) {
		return
	}

	s.vars = append(s.vars, v)
}

/* if variable exists, override variable
 * else add the variable
 */
func (s *Section) override(v Tuple_v) {
	for i, _ := range s.vars {
		if (s.vars[i].name) == v.name {
			s.vars[i].value = v.value
			return
		}
	}
	s.insert(v)
}

func NewSection(name string) Section {
	return Section{
		name: name,
	}
}

type Object struct {
	sections []Section
	current  int
}

func (o Object) has(key string) bool {
	for _, s := range o.sections {
		if s.name == key {
			return true
		}
	}
	return false
}

func (o Object) at(key string) Section {
	for _, s := range o.sections {
		if s.name == key {
			return s
		}
	}
	return o.sections[0]
}

/* if section exists, exit
 * else push_back section and set to currentsection
 */
func (o *Object) insert(name string) {
	if o.has(name) {
		return
	}
	o.sections = append(o.sections, NewSection(name))
	o.current = len(o.sections) - 1
}

/* if section exists, set to currentsection
 * else push_back section and set to currentsection
 */
func (o *Object) override(name string) {
	for i, s := range o.sections {
		if s.name == name {
			o.current = i
			return
		}
	}
	o.insert(name)
}

func (o *Object) Add(section string, key string, val string) {
	o.override(section)
	o.sections[o.current].override(Tuple_v{ key, val })
}

func (o Object) Get(section string, key string) string {
	s := o.at(section)
	v := s.at(key)
	return v.value
}

func NewObject() Object {

	s := NewSection("default")

	return Object{
		sections: []Section{s},
		current:  0,
	}
}

type Ini struct {
	format Format
	errors []string

	Object Object
}

func (self *Ini) Parse(in bufio.Scanner) {
	var line string

	for in.Scan() {
		line = in.Text()

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			self.errors = append(self.errors, line)
			continue
		}

		// pos := strings.Index(line, self.format.char_assign)
		pos := func(line string, sub string) int {
			for i, s := range(line) {
				if string( s ) == sub {
					return i
				}
			}
			return -1
		}(line, self.format.char_assign)
		front := string(line[0])
		end := string(line[len(line)-1])

		if self.format.is_comment(front) {
			continue
		}
		if self.format.is_section_start(front) {
			if self.format.is_section_end(end) {
				self.Object.override(string(line[1 : len(line)-1]))
			} else {
				self.errors = append(self.errors, line)
			}
			continue
		}

		if pos > 0 && pos != len(line)-1 {
			name := strings.TrimSpace(line[0:pos])
			val := strings.TrimSpace(line[pos+1:])

			self.Object.sections[self.Object.current].override(Tuple_v{name, val})
		}
	}
}

func NewIni() Ini {
	return Ini{
		format: NewFormat(),
		Object: NewObject(),
	}
}

func (s Section) ToString() string {
	var str string

	str += "[" + s.name + "]"
	
	for _, v := range(s.vars) {
		str += v.name + "=" + v.value
	}

	return str
}

func (o Object) ToString() string {
	var str string

	for _, s := range(o.sections) {
		str += s.ToString()
	}

	return str
}
