package model

import (
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/hashicorp/errwrap"
)

type Model struct {
	db     *bolt.DB
	DBPath string

	Events chan Event
}

func NewModel(dbpath string) (*Model, error) {
	opts := &bolt.Options{Timeout: time.Second}

	db, err := bolt.Open(dbpath, 0600, opts)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Failed to open Dockpit db at '%s': {{err}}, is Dockpit already running?", dbpath), err)
	}

	m := &Model{db: db, DBPath: dbpath, Events: make(chan Event)}
	err = m.UpsertMetaData()
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Failed to initialize metadata for '%s': {{err}}", dbpath), err)
	}

	return m, nil
}

func (m *Model) GetDBMetaData() (*Meta, error) {
	var meta *Meta

	return meta, m.db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(MetaBucketName))
		if b == nil {
			return fmt.Errorf("Failed to open meta bucket")
		}

		data := b.Get([]byte(DatabaseMetaKey))
		if data != nil {
			meta, err = NewMetaFromSerialized(data)
			if err != nil {
				return errwrap.Wrapf(fmt.Sprintf("Failed to deserialize meta from db: {{err}}, data: '%s'", string(data)), err)
			}
		}

		return nil
	})
}

func (m *Model) UpsertMetaData() error {
	return m.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(MetaBucketName))
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to create meta bucket: {{err}}"), err)
		}

		data := b.Get([]byte(DatabaseMetaKey))
		var dbmeta *Meta
		if data == nil {
			dbmeta, err = NewMeta()
			if err != nil {
				return errwrap.Wrapf(fmt.Sprintf("Failed to create new db meta: {{err}}"), err)
			}
		} else {
			dbmeta, err = NewMetaFromSerialized(data)
			if err != nil {
				return errwrap.Wrapf(fmt.Sprintf("Failed to deserialize db meta data '%s': {{err}}", string(data)), err)
			}
		}

		data, err = dbmeta.Serialize()
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to serialize db meta '%s': {{err}}", dbmeta.ID), err)
		}

		b.Put([]byte(DatabaseMetaKey), data)
		return nil
	})
}

func (m *Model) GetAllIsolations() ([]*Isolation, error) {
	isos := []*Isolation{}

	err := m.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(IsolationBucketName))
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to create isolation bucket: {{err}}"), err)
		}

		return b.ForEach(func(k, v []byte) error {
			iso, err := NewIsolationFromSerialized(v)
			if err != nil {
				return errwrap.Wrapf(fmt.Sprintf("Failed to deserialize isolation from db: {{err}}, data: '%s'", string(v)), err)
			}

			isos = append(isos, iso)
			return nil
		})
	})

	return isos, err
}

func (m *Model) GetAllDeps() ([]*Dep, error) {
	deps := []*Dep{}

	err := m.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(DepBucketName))
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to create deps bucket: {{err}}"), err)
		}

		return b.ForEach(func(k, v []byte) error {
			dep, err := NewDepFromSerialized(v)
			if err != nil {
				return errwrap.Wrapf(fmt.Sprintf("Failed to deserialize dep from db: {{err}}, data: '%s'", string(v)), err)
			}

			deps = append(deps, dep)
			return nil
		})
	})

	return deps, err
}

func (m *Model) GetIsolationDepsAndStates(iso *Isolation) (map[*Dep]*State, error) {
	res := map[*Dep]*State{}

	for dname, sname := range iso.States {
		dep, err := m.FindDepByName(dname)
		if err != nil {
			return res, errwrap.Wrapf(fmt.Sprintf("Failed to find dep with name '%s': {{err}}", dname), err)
		}

		s := dep.GetState(sname)
		if s == nil {
			return res, fmt.Errorf("Dependency '%s' doesn't have a state with name '%s'", dep.Name, sname)
		}

		res[dep] = dep.GetState(sname)
	}

	return res, nil
}

func (m *Model) FindDepByName(name string) (*Dep, error) {
	var dep *Dep
	return dep, m.db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(DepBucketName))
		if b == nil {
			return fmt.Errorf("Failed to open dep bucket")
		}

		data := b.Get([]byte(name))
		if data != nil {
			dep, err = NewDepFromSerialized(data)
			if err != nil {
				return errwrap.Wrapf(fmt.Sprintf("Failed to deserialize isolation from db: {{err}}, data: '%s'", string(data)), err)
			}
		}

		return nil
	})
}

func (m *Model) FindIsolationByID(ID string) (*Isolation, error) {
	var iso *Isolation
	return iso, m.db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(IsolationBucketName))
		if b == nil {
			return fmt.Errorf("Failed to open isolation bucket")
		}

		data := b.Get([]byte(ID))
		if data != nil {
			iso, err = NewIsolationFromSerialized(data)
			if err != nil {
				return errwrap.Wrapf(fmt.Sprintf("Failed to deserialize isolation from db: {{err}}, data: '%s'", string(data)), err)
			}
		}

		return nil
	})
}

func (m *Model) FindIsolationByName(name string) (*Isolation, error) {
	var iso *Isolation
	return iso, m.db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(IsolationBucketName))
		if b == nil {
			return fmt.Errorf("Failed to open isolation bucket")
		}

		data := b.Get([]byte(name))
		if data != nil {
			iso, err = NewIsolationFromSerialized(data)
			if err != nil {
				return errwrap.Wrapf(fmt.Sprintf("Failed to deserialize isolation from db: {{err}}, data: '%s'", string(data)), err)
			}
		}

		return nil
	})
}

func (m *Model) RemoveDepStateByID(dep *Dep, sid string) error {
	for i, state := range dep.States {
		if state.ID == sid {

			//remove any deps from iso
			isos, err := m.GetAllIsolations()
			if err != nil {
				return errwrap.Wrapf(fmt.Sprintf("Failed to get all isolations while removing dep state with id '%s::%s': {{err}}", dep.Name, sid), err)
			}

			for _, iso := range isos {
				if iso.HasDepState(dep, sid) {
					iso.RemoveDep(dep)
					err = m.UpdateIsolation(iso)
					if err != nil {
						return errwrap.Wrapf(fmt.Sprintf("Could not update isolation '%s' with removed dep '%s': {{err}}", iso.Name, dep.Name), err)
					}
				}
			}

			dep.States = append(dep.States[:i], dep.States[i+1:]...)
		}
	}

	err := m.UpdateDep(dep)
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed update dep with id '%s': {{err}}", sid), err)
	}

	return nil
}

func (m *Model) RemoveDep(dep *Dep) error {

	//first, remove dep from any isolation
	isos, err := m.GetAllIsolations()
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Failed to get all isolations while removing dep with name '%s': {{err}}", dep.Name), err)
	}

	for _, iso := range isos {
		if iso.HasDep(dep) {
			iso.RemoveDep(dep)
			err = m.UpdateIsolation(iso)
			if err != nil {
				return errwrap.Wrapf(fmt.Sprintf("Could not update isolation '%s' with removed dep '%s': {{err}}", iso.Name, dep.Name), err)
			}
		}
	}

	//second, remove dep itself
	err = m.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(DepBucketName))
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to create dep bucket: {{err}}"), err)
		}

		err = b.Delete([]byte(dep.Name))
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to remove dep with name '%s': {{err}}", dep.Name), err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	m.Events <- NewEvent(ActionRemoved, dep)
	return nil
}

func (m *Model) RemoveIsolation(iso *Isolation) error {
	err := m.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(IsolationBucketName))
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to create isolation bucket: {{err}}"), err)
		}

		err = b.Delete([]byte(iso.ID))
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to remove isolation with name '%s': {{err}}", iso.Name), err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	m.Events <- NewEvent(ActionRemoved, iso)
	return nil
}

func (m *Model) InsertDep(dep *Dep) error {
	err := m.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(DepBucketName))
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to create dep bucket: {{err}}"), err)
		}

		existing := b.Get([]byte(dep.Name))
		if existing != nil {
			return fmt.Errorf("A dependency with name '%s' already exists", dep.Name)
		}

		data, err := dep.Serialize()
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to serialize dep '%s': {{err}}", dep.Name), err)
		}

		err = b.Put([]byte(dep.Name), data)
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to insert serialized isolation '%s': {{err}}", dep.Name), err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	m.Events <- NewEvent(ActionCreated, dep)
	return nil
}

func (m *Model) UpdateDep(dep *Dep) error {
	err := m.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(DepBucketName))
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to create Dep bucket: {{err}}"), err)
		}

		data, err := dep.Serialize()
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to serialize Dep '%s': {{err}}", dep.Name), err)
		}

		err = b.Put([]byte(dep.Name), data)
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to update serialized Dep '%s': {{err}}", dep.Name), err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	m.Events <- NewEvent(ActionUpdate, dep)
	return nil
}

func (m *Model) UpdateIsolation(iso *Isolation) error {
	err := m.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(IsolationBucketName))
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to create isolation bucket: {{err}}"), err)
		}

		data, err := iso.Serialize()
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to serialize isolation '%s': {{err}}", iso.Name), err)
		}

		err = b.Put([]byte(iso.ID), data)
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to update serialized isolation '%s': {{err}}", iso.Name), err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	m.Events <- NewEvent(ActionUpdate, iso)
	return nil
}

func (m *Model) InsertIsolation(iso *Isolation) error {
	err := m.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(IsolationBucketName))
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to create isolation bucket: {{err}}"), err)
		}

		data, err := iso.Serialize()
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to serialize isolation '%s': {{err}}", iso.Name), err)
		}

		err = b.Put([]byte(iso.ID), data)
		if err != nil {
			return errwrap.Wrapf(fmt.Sprintf("Failed to insert serialized isolation '%s': {{err}}", iso.Name), err)
		}

		return nil
	})

	if err != nil {
		return err
	}

	m.Events <- NewEvent(ActionCreated, iso)
	return nil
}
